// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import Api from '@/api/Api'

require('semantic-ui-css/semantic.css')
require('semantic-ui-css/semantic.js')

Vue.config.productionTip = false

Vue.mixin({
  data: function () {
    return {
      get api () {
        return new Api()
      }
    }
  }
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})
