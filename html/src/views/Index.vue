<template>
    <div class="libra-index">
        <div class="libra-cards">
            <Card title="hah"></Card>
            <Card v-for="count in 90"
                  :key="count"
                  :img="`http://www.thiswaifudoesnotexist.net/example-${count * 1000}.jpg`"
                  title="测试测试"
                  to="http://www.baidu.com">
            </Card>
            <Menu></Menu>
        </div>
    </div>
</template>

<script>

    import Card from '@/components/Card';
    import Menu from '@/components/Menu';

  export default {
    name: "Index",
    components:{
      Card,
      Menu
    }
  };
</script>
