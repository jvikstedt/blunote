<template>
  <div>
    <div class="action">
      <search :onChange="onSearchChange"></search>
    </div>
    <div class="result">
      <note-list :notes="notes" :onClick="onNoteClick"></note-list>
    </div>
    <router-link to="/notes/new" class="ui button">New</router-link>
  </div>
</template>

<script>
import { EventBus } from '@/EventBus'
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
        EventBus.$emit('flash', { status: 'error', header: 'Something went wrong!', body: e.toString() })
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
      EventBus.$emit('flash', { status: 'error', header: 'Something went wrong!', body: e.toString() })
    }
  }
}
</script>

<style>
  .result {
    margin: 1em 0;
  }
</style>
