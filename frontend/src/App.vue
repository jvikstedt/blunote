<template>
  <div id="app">
    <div class="ui fixed inverted menu">
      <div class="ui container">
        <router-link to="/" class="header item">
          dsNotes
        </router-link>
        <router-link to="/" class="item">Home</router-link>
      </div>
    </div>

    <div class="ui main text container">
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
    background-color: #FFFFFF;
  }
  .ui.menu .item img.logo {
    margin-right: 1.5em;
  }
  .main.container {
    margin-top: 5em;
  }
</style>
