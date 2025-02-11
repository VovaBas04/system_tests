<template>
  <HeaderCodeChecker/>
  <SendTaskDescription @send="sendTask"/>
  <SelectView/>
  <ModalWindow model-value title="Отправить код" :hidden="isHidden">
    <CodeEditor/>
  </ModalWindow>
  <router-view></router-view>
</template>

<script>

import HeaderCodeChecker from "@/components/HeaderCodeChecker.vue";
import ModalWindow from "@/components/ModalWindow.vue";
import CodeEditor from "@/components/CodeEditor.vue";
import SendTaskDescription from "@/views/SendTaskDescription.vue";
import SelectView from "@/views/SelectView.vue";

export default {
  name: 'App',
  components: {
    SelectView,
    SendTaskDescription,
    CodeEditor,
    ModalWindow,
    HeaderCodeChecker,
  },
  data() {
    return {
      hidden: true
    }
  },
  methods: {
    sendTask() {
      this.hidden = false
    }
  },
  computed: {
    isHidden() {
      return this.hidden
    }
  }
}
</script>

<style>
</style>
