<template>
    <div class="libra-card" @click="toHandler">
        <div class="libra-card-category">
            测试
        </div>
        <div class="libra-card-icon">
            <img :src="value" v-if="value">
            <div class="libra-card-word" v-else>
                {{title.substring(0,1)}}
            </div>
        </div>
        <div class="libra-card-title">
            {{title}}
        </div>
    </div>
</template>

<script>

  import util from '@/util/util';

  export default {
    name: "Card",
    props:{
      img:{
        type: String,
        default: null
      },
      to: {
        type: String,
        default: null
      },
      title: {
        type: String,
        default: null,
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
      },
      toHandler(){
        if (this.to) {
          window.open(this.to,"_blank")
        }
      }
    }
  };
</script>


