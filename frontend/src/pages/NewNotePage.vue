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

  methods: {
    onFormSubmit: async function (e) {
      try {
        await this.api.post(`/notes`, this.note)
        EventBus.$emit('flash', { status: 'success', header: 'Created successfully!' })
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
