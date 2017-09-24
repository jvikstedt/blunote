<template>
  <note-list v-bind:notes="notes"></note-list>
</template>

<script type="text/javascript">
import NoteList from '../components/NoteList'

export default {
  data: () => ({
    notes: []
  }),

  components: {
    NoteList
  },

  async created () {
    try {
      this.notes = await this.api.get('/notes')
    } catch (e) {
      console.log(e)
    }
  }
}
</script>

<style>
</style>
