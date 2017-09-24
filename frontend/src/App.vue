<template>
  <div id="app">
    <div class="wrapper">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
export default {
  name: 'app'
}
</script>

<style>
  body {
    padding: 1em;
  }

  .wrapper {
    max-width: 960px;
    margin: 20px auto;
  }
</style>
