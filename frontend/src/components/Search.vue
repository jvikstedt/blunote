<template>
  <div class="ui input focus">
    <input type="text" placeholder="Search..." v-model="userInput">
  </div>
</template>

<script type="text/javascript">
export default {
  props: ['onChange'],
  data: () => ({
    userInput: ''
  }),
  watch: {
    userInput: function (val, oldVal) {
      this.onChange(val)
    }
  }
}
</script>
<style>
</style>
