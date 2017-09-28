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
import quill from '../quill'

export default {
  data: () => ({
    note: {},
    content: '',
    editorOption: quill.options
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
  }
}
</script>

<style>
</style>
