<template>
  <div>
    {{ note.title }}
    {{ note.body }}
  </div>
</template>

<script type="text/javascript">

export default {
  data: () => ({
    note: {}
  }),

  async created () {
    const noteId = this.$route.params.noteId
    try {
      this.note = await this.api.get(`/notes/${noteId}`)
    } catch (e) {
      console.log(e)
    }
  }
}
</script>

<style>
</style>
