<template>
  <div>
    <input name="title" type="text" v-model="note.title" />
    <quill-editor v-model="note.body_html"
      :options="editorOption"
      ref="editor">
    </quill-editor>
    <button class="ui button" @click="onSubmit">Submit</button>
  </div>
</template>

<script type="text/javascript">
import { quillEditor } from 'vue-quill-editor'
import quill from '../quill'

export default {
  props: {
    note: {
      type: Object,
      default () { return {} }
    },
    onSubmit: {
      type: Function,
      required: true
    }
  },
  data: () => ({
    editorOption: quill.options
  }),

  components: {
    quillEditor
  }
}
</script>

<style>
</style>
