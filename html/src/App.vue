<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<style lang="scss">
  * {
    padding: 0;
    margin: 0;
  }
  body{
    height: 100vh;
    width: 100vw;
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    color: #42b983;
    overflow-x: hidden;
  }
  #app{
    width: 100%;
    height: 100%;
  }
</style>
