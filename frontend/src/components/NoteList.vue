<template>
  <div>
    <div class="ui card" v-for="(item, index) in notes" @click="onClick(item)">
      <div class="content">
        <div class="header">{{ item.title }}</div>
        <div class="description"><p>{{ item.body }}</p></div>
      </div>
    </div>
  </div>
</template>

<script type="text/javascript">
export default {
  props: ['notes', 'onClick']
}
</script>
<style>
  .ui.card {
    width: 100%;
    cursor: pointer;
  }

  .ui.card:hover {
    background-color: #f8f8f8;
  }
</style>
