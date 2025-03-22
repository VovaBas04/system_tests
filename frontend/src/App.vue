<template>
  <HeaderCodeChecker/>
  <router-view></router-view>
</template>

<script>

import HeaderCodeChecker from "@/components/HeaderCodeChecker.vue";

export default {
  name: 'App',
  components: {
    HeaderCodeChecker,
  }
}
</script>

<style>
</style>
