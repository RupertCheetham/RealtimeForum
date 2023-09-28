export default class AbstractView {
	constructor() {}

	setTitle(title) {
		document.title = title
	}

	async getHTML() {
		return ""
	}
}
