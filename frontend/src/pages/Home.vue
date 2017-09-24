<template>
  <div>
    <div class="action">
      <search :onChange="onSearchChange"></search>
    </div>
    <div class="result">
      <note-list :notes="notes" :onClick="onNoteClick"></note-list>
    </div>
  </div>
</template>

<script type="text/javascript">
import _ from 'lodash'

import NoteList from '../components/NoteList'
import Search from '../components/Search'

export default {
  data: () => ({
    notes: []
  }),

  components: {
    NoteList,
    Search
  },

  methods: {
    onSearchChange: _.throttle(async function (input) {
      try {
        this.notes = await this.api.get(`/notes?search=${input}`)
      } catch (e) {
        console.log(e)
      }
    }, 300),
    onNoteClick: function (note) {
      this.$router.push({ path: `/notes/${note.id}` })
    }
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
  .result {
    margin-top: 1em;
  }
</style>
