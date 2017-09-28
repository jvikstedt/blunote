<template>
  <div>
    {{ note.title }}
    {{ note.body }}
    <div>
      <quill-editor v-model="content"
        ref="editor">
      </quill-editor>
    </div>
    <button @click="onUpdateNote">Update</button>
  </div>
</template>

<script type="text/javascript">
import { quillEditor } from 'vue-quill-editor'

export default {
  data: () => ({
    note: {},
    content: ''
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
