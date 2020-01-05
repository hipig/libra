<template>
    <div class="libra-bg">
        <img class="libra-img" :src="value" v-if="value">
    </div>
</template>

<script>

  import util from '@/util/util';

  export default {
    name: "BG",
    props:{
      img:{
        type: String,
        default: null
      }
    },
    data(){
      return {
        value: null
      }
    },
    async mounted(){
      this.loadImg()
    },
    methods:{
      loadImg(){
        if(this.img){
          let _this = this
          util.loadImageAsync(this.img).then(v =>{
            _this.value = v.src
          })
        }
      }
    }
  };
</script>
