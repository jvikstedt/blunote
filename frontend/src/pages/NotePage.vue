<template>
  <div>
    <div class="ui header teal">{{ note.title }}</div>
    <div class="ui stacked segment" v-html="note.body_html" />
    <div class="ui">
      <router-link :to="'/notes/' + note.id + '/edit'" class="ui button">Edit</router-link>
    </div>
  </div>
</template>

<script type="text/javascript">
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
  }
}
</script>

<style>
</style>
