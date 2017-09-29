<template>
  <div>
    <NoteForm :note="note" :onSubmit="onFormSubmit" />
  </div>
</template>

<script>
import NoteForm from '../components/NoteForm'
import { EventBus } from '@/EventBus'

export default {
  data: () => ({
    note: {}
  }),

  async created () {
    const noteId = this.$route.params.noteId
    try {
      this.note = await this.api.get(`/notes/${noteId}`)
    } catch (e) {
      EventBus.$emit('flash', { status: 'error', header: 'Something went wrong!', body: e.toString() })
    }
  },

  methods: {
    onFormSubmit: async function (e) {
      try {
        await this.api.put(`/notes/${this.note.id}`, this.note)
        EventBus.$emit('flash', { status: 'success', header: 'Updated successfully!' })
      } catch (e) {
        EventBus.$emit('flash', { status: 'error', header: 'Something went wrong!', body: e.toString() })
      }
    }
  },

  components: {
    NoteForm
  }
}
</script>

<style>
</style>
