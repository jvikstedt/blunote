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
import { EventBus } from '@/EventBus'

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
      EventBus.$emit('flash', { status: 'error', header: 'Something went wrong!', body: e.toString() })
    }
  },

  methods: {
    onUpdateNote: async function () {
      try {
        this.note = await this.api.put(`/notes/${this.note.id}`, {
          ...this.note,
          body_html: this.content
        })
        EventBus.$emit('flash', { status: 'success', header: 'Updated successfully!' })
      } catch (e) {
        EventBus.$emit('flash', { status: 'error', header: 'Something went wrong!', body: e.toString() })
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
