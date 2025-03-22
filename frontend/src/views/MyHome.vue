<template>
  <HeaderCodeChecker/>
    <SendTaskDescription @send="sendTask" @update:tests="setTests"/>
    <ModalWindow model-value title="Отправить код" :hidden="isHidden" :tests="tests"/>
  <router-view></router-view>
</template>

<script>

import HeaderCodeChecker from "@/components/HeaderCodeChecker.vue";
import ModalWindow from "@/components/ModalWindow.vue";
import SendTaskDescription from "@/views/SendTaskDescription.vue";

export default {
  name: 'App',
  components: {
    SendTaskDescription,
    ModalWindow,
    HeaderCodeChecker,
  },
  data() {
    return {
      hidden: true,
      tests: [],
    }
  },
  methods: {
    sendTask() {
      this.hidden = false
    },
    setTests(tests) {
      console.log("Тут", tests)
      this.tests = tests
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
