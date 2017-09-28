
function imageHandler () {
  var range = this.quill.getSelection()
  var value = prompt('What is the image URL')
  this.quill.insertEmbed(range.index, 'image', value, 'user')
}

export default {
  image: imageHandler
}
