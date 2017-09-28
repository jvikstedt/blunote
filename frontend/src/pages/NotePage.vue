<template>
  <div>
    <h1>
      {{ note.title }}
    </h1>
    <div>
      <quill-editor v-model="content"
        :options="editorOption"
        ref="editor">
      </quill-editor>
    </div>
    <button class="ui button" @click="onUpdateNote">Update</button>
  </div>
</template>

<script type="text/javascript">
import { quillEditor } from 'vue-quill-editor'

function imageHandler () {
  var range = this.quill.getSelection()
  var value = prompt('What is the image URL')
  this.quill.insertEmbed(range.index, 'image', value, 'user')
}

export default {
  data: () => ({
    note: {},
    content: '',
    editorOption: {
      modules: {
        toolbar: {
          container: [
            ['bold', 'italic', 'underline', 'strike'],
            ['blockquote', 'code-block'],
            [{ 'header': 1 }, { 'header': 2 }],
            [{ 'list': 'ordered' }, { 'list': 'bullet' }],
            [{ 'script': 'sub' }, { 'script': 'super' }],
            [{ 'indent': '-1' }, { 'indent': '+1' }],
            [{ 'direction': 'rtl' }],
            [{ 'size': ['small', false, 'large', 'huge'] }],
            [{ 'header': [1, 2, 3, 4, 5, 6, false] }],
            [{ 'color': [] }, { 'background': [] }],
            [{ 'font': [] }],
            [{ 'align': [] }],
            ['clean'],
            ['link', 'image', 'video']
          ],
          handlers: {
            image: imageHandler
          }
        }
      }
    }
  }),

  async created () {
    const noteId = this.$route.params.noteId
    try {
      this.note = await this.api.get(`/notes/${noteId}`)
      this.content = this.note.body_html
    } catch (e) {
      console.log(e)
    }
  },

  methods: {
    onUpdateNote: async function () {
      try {
        this.note = await this.api.put(`/notes/${this.note.id}`, {
          ...this.note,
          body_html: this.content
        })
      } catch (e) {
        console.log(e)
      }
    }
  },

  components: {
    quillEditor
  },

  computed: {
    editor () {
      return this.$refs.editor.quill
    }
  }
}
</script>

<style>
</style>
