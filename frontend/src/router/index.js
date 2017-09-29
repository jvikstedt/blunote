import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/pages/Home'
import NotePage from '@/pages/NotePage'
import NewNotePage from '@/pages/NewNotePage'
import EditNotePage from '@/pages/EditNotePage'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/notes/new',
      name: 'newNote',
      component: NewNotePage
    },
    {
      path: '/notes/:noteId',
      name: 'note',
      component: NotePage
    },
    {
      path: '/notes/:noteId/edit',
      name: 'editNote',
      component: EditNotePage
    }
  ]
})
