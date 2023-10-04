export default class AbstractView {
	constructor() {}

	setTitle(title) {
		document.title = title
	}

	async renderHTML() {
		return ""
	}
}
