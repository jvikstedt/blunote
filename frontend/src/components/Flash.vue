<template>
  <transition name="fade">
    <div v-if="current" class="flash-message">
      <div :class="'ui message ' + current.status">
        <i class="close icon" @click="removeCurrent"></i>
        <div class="header">
          {{ current.header }}
        </div>
        <p v-if="current.body">{{ current.body }}</p>
      </div>
    </div>
  </transition>
</template>

<script type="text/javascript">
import { EventBus } from '@/EventBus'

export default {
  data: () => ({
    current: false
  }),

  created: function () {
    EventBus.$on('flash', event => {
      this.current = event
    })
  },

  methods: {
    removeCurrent: function (event) {
      this.current = false
    }
  }
}
</script>
<style>
  .flash-message {
    margin-bottom: 1em;
  }
  .fade-enter-active, .fade-leave-active {
    transition: opacity .5s
  }
  .fade-enter, .fade-leave-to {
    opacity: 0
  }
</style>
