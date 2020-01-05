<template>
  <div id="app">
    <BG img="https://cdn.jsdelivr.net/gh/lhlyu/pb@master/2019/6.jpg"></BG>
    <router-view />
  </div>
</template>

<script>

  import BG from '@/components/BG';

  export default {
    name: "Index",
    components:{
      BG
    }
  };
</script>

<style lang="scss">
@import "assets/style/style";
</style>
