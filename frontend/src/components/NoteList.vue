<template>
  <div class="ui list">
    <div class="item" v-for="(item, index) in notes">
      <router-link :to="'/notes/' + item.id">
        <div class="header">{{ item.title }}</div>
        {{ item.body }}
      </router-link>
    </div>
  </div>
</template>

<script type="text/javascript">
export default {
  props: ['notes']
}
</script>

<style>
</style>
