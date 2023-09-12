fetch('http://localhost:8080/posts')
  .then(response => response.json())
  .then(data => {
    // Handle the data here
    console.log(data);
  })
  .catch(error => {
    console.error('Error fetching data:', error);
  });
