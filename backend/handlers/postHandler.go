package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"realtimeForum/db"
	"time"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB
// Handler for posts page
func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	if r.Method == "POST" {
		var post db.PostEntry
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// 32 MB is the default used by FormFile()
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Get a reference to the fileHeaders.
		// They are accessible only after ParseMultipartForm is called
		files := r.MultipartForm.File["file"]

		for _, fileHeader := range files {
			// Restrict the size of each uploaded file to 1MB.
			// To prevent the aggregate size from exceeding
			// a specified value, use the http.MaxBytesReader() method
			// before calling ParseMultipartForm()
			if fileHeader.Size > MAX_UPLOAD_SIZE {
				http.Error(w, fmt.Sprintf("The uploaded image is too big: %s. Please use an image less than 1MB in size", fileHeader.Filename), http.StatusBadRequest)
				return
			}

			// Open the file
			file, err := fileHeader.Open()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			defer file.Close()

			buff := make([]byte, 512)
			_, err = file.Read(buff)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			filetype := http.DetectContentType(buff)
			if filetype != "image/jpeg" && filetype != "image/png" {
				http.Error(w, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
				return
			}

			_, err = file.Seek(0, io.SeekStart)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = os.MkdirAll("./uploads", os.ModePerm)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			f, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			defer f.Close()

			_, err = io.Copy(f, file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		fmt.Fprintf(w, "Upload successful")

		fmt.Println("post is:", post)

		// utils.WriteMessageToLogFile("Received post: " + "post.Body")
		log.Println("Received post:", post.Username, post.Img, post.Body, post.Categories)

		err = db.AddPostToDatabase(post.Username, post.Img, post.Body, post.Categories)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}

	if r.Method == "GET" {
		posts, err := db.GetPostFromDatabase()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(posts) > 0 {
			json.NewEncoder(w).Encode(posts)
		} else {
			w.Write([]byte("No posts available"))
		}
	}

}

// type Progress struct {
// 	TotalSize int64
// 	BytesRead int64
// }

// // Write is used to satisfy the io.Writer interface.
// // Instead of writing somewhere, it simply aggregates
// // the total bytes on each read
// func (pr *Progress) Write(p []byte) (n int, err error) {
// 	n, err = len(p), nil
// 	pr.BytesRead += int64(n)
// 	pr.Print()
// 	return
// }

// // Print displays the current progress of the file upload
// // each time Write is called
// func (pr *Progress) Print() {
// 	if pr.BytesRead == pr.TotalSize {
// 		fmt.Println("DONE!")
// 		return
// 	}

// 	fmt.Printf("File upload in progress: %d\n", pr.BytesRead)
// }

// const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB

// func uploadHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// 32 MB is the default used by FormFile()
// 	if err := r.ParseMultipartForm(32 << 20); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Get a reference to the fileHeaders.
// 	// They are accessible only after ParseMultipartForm is called
// 	files := r.MultipartForm.File["file"]

// 	for _, fileHeader := range files {
// 		// Restrict the size of each uploaded file to 1MB.
// 		// To prevent the aggregate size from exceeding
// 		// a specified value, use the http.MaxBytesReader() method
// 		// before calling ParseMultipartForm()
// 		if fileHeader.Size > MAX_UPLOAD_SIZE {
// 			http.Error(w, fmt.Sprintf("The uploaded image is too big: %s. Please use an image less than 1MB in size", fileHeader.Filename), http.StatusBadRequest)
// 			return
// 		}

// 		// Open the file
// 		file, err := fileHeader.Open()
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		defer file.Close()

// 		buff := make([]byte, 512)
// 		_, err = file.Read(buff)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		filetype := http.DetectContentType(buff)
// 		if filetype != "image/jpeg" && filetype != "image/png" {
// 			http.Error(w, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
// 			return
// 		}

// 		_, err = file.Seek(0, io.SeekStart)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		err = os.MkdirAll("./uploads", os.ModePerm)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		f, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		defer f.Close()

// 		_, err = io.Copy(f, file)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
// 	}

// 	fmt.Fprintf(w, "Upload successful")
// }

// func UploadHandler(w http.ResponseWriter, r *http.Request, file multipart.File, fileHeader *multipart.FileHeader) string {
// 	err := os.MkdirAll("./static/uploads", os.ModePerm)
// 	if err != nil {
// 		utils.HandleError("error creating file directory for uploads", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return ""
// 	}

// 	destinationFile, err := os.Create(fmt.Sprintf("./static/uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
// 	if err != nil {
// 		utils.HandleError("error creating file for image", err)
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return ""
// 	}

// 	defer destinationFile.Close()

// 	_, err = io.Copy(destinationFile, file)

// 	if err != nil {
// 		utils.HandleError("error copying file to destination", err)
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return ""
// 	}

// 	utils.WriteMessageToLogFile("file uploaded successfully")
// 	fileName := destinationFile.Name()[1:]
// 	return fileName
// }

// // followed this: https://freshman.tech/file-upload-golang/
// func SubmissionHandler(w http.ResponseWriter, r *http.Request, user dbmanagement.User, tmpl *template.Template) {
// 	// 20 megabytes
// 	idToDelete := r.FormValue("deletepost")
// 	if idToDelete != "" {
// 		dbmanagement.DeletePostWithUUID(idToDelete)
// 	}
// 	notificationToDelete := r.FormValue("delete notification")
// 	if notificationToDelete != "" {
// 		dbmanagement.DeleteFromTableWithUUID("Notifications", notificationToDelete)
// 	}

// 	like := r.FormValue("like")
// 	dislike := r.FormValue("dislike")

// 	if like != "" {
// 		dbmanagement.AddReactionToPost(user.UUID, like, 1)
// 		post, err := dbmanagement.SelectPostFromUUID(like)
// 		if err != nil {
// 			PageErrors(w, r, tmpl, 500, "Internal Server Error")
// 			return
// 		}
// 		receiverId, _ := dbmanagement.SelectUserFromName(post.OwnerId)
// 		dbmanagement.AddNotification(receiverId.UUID, like, "", user.UUID, 1, "")
// 	}
// 	if dislike != "" {
// 		dbmanagement.AddReactionToPost(user.UUID, dislike, -1)
// 		post, err := dbmanagement.SelectPostFromUUID(dislike)
// 		if err != nil {
// 			PageErrors(w, r, tmpl, 500, "Internal Server Error")
// 		}
// 		receiverId, _ := dbmanagement.SelectUserFromName(post.OwnerId)
// 		dbmanagement.AddNotification(receiverId.UUID, dislike, "", user.UUID, -1, "")
// 	}

// 	maxSize := 20 * 1024 * 1024

// 	r.Body = http.MaxBytesReader(w, r.Body, int64(maxSize))
// 	err := r.ParseMultipartForm(int64(maxSize))
// 	if err != nil {
// 		// only actual post submissions have multipart enabled, deleting, likes, dislikes aren't mulipart but that's already handled above so can end function
// 		if err.Error() == "request Content-Type isn't multipart/form-data" {
// 			return
// 		}
// 		utils.HandleError("error parsing form for image, likely too big", err)
// 		return
// 	}

// 	file, fileHeader, err := r.FormFile("submission-image")
// 	fileName := ""
// 	if err != nil {
// 		// if you were trying to make a post without an image it will log this 'error' but still submit the text and tags
// 		utils.HandleError("error retrieving file from form", err)
// 	} else {
// 		utils.WriteMessageToLogFile("trying to retrieve file...")
// 		defer file.Close()
// 		fileName = UploadHandler(w, r, file, fileHeader)
// 	}

// 	title := r.FormValue("submission-title")
// 	content := r.FormValue("post")
// 	tags := r.Form["tags"]
// 	edit := r.FormValue("editpost")

// 	if edit != "" {
// 		if CheckInputs(content) && CheckInputs(title) {
// 			userFromUUID, err := dbmanagement.SelectUserFromUUID(user.UUID)
// 			utils.HandleError("cant get user with uuid in all posts", err)
// 			getLikes, err := dbmanagement.SelectPostFromUUID(edit)
// 			if err != nil {
// 				PageErrors(w, r, tmpl, 500, "Internal Server Error")
// 			}
// 			getDislikes, err := dbmanagement.SelectPostFromUUID(edit)
// 			if err != nil {
// 				PageErrors(w, r, tmpl, 500, "Internal Server Error")
// 			}
// 			editedPost, err := dbmanagement.UpdatePost(edit, title, content, userFromUUID.Name, getLikes.Likes, getDislikes.Dislikes, time.Now(), fileName)
// 			if err != nil {
// 				PageErrors(w, r, tmpl, 500, "Internal Server Error")
// 			}
// 			dbmanagement.UpdateTaggedPost(edit)
// 			for _, tag := range tags {
// 				InputTags(tag, editedPost)
// 				InputTags(tag, editedPost)
// 			}

// 		}
// 	} else {
// 		if CheckInputs(content) && CheckInputs(title) {
// 			userFromUUID, err := dbmanagement.SelectUserFromUUID(user.UUID)
// 			utils.HandleError("cant get user with uuid in all posts", err)
// 			post, err := dbmanagement.InsertPost(title, content, userFromUUID.Name, 0, 0, time.Now(), fileName)
// 			if err != nil {
// 				PageErrors(w, r, tmpl, 500, "Internal Server Error")
// 			}
// 			for _, tag := range tags {
// 				InputTags(tag, post)
// 				InputTags(tag, post)
// 			}
// 		}
// 	}
// }
