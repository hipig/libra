<template>
    <div class="layout">
        <Layout>
            <Sider ref="side1" hide-trigger collapsible :collapsed-width="78" v-model="isCollapsed">
                <Menu active-name="1-2" theme="dark" width="auto" :class="menuitemClasses">
                    <h1>Admin</h1>
                    <MenuGroup title="内容管理">
                        <MenuItem name="1">
                            <Icon type="md-document" />
                            文章管理
                        </MenuItem>
                        <MenuItem name="2">
                            <Icon type="md-chatbubbles" />
                            评论管理
                        </MenuItem>
                    </MenuGroup>
                    <MenuGroup title="统计分析">
                        <MenuItem name="3">
                            <Icon type="md-heart" />
                            用户留存
                        </MenuItem>
                        <MenuItem name="4">
                            <Icon type="md-leaf" />
                            流失用户
                        </MenuItem>
                    </MenuGroup>
                </Menu>
            </Sider>
            <Layout>
                <Header :style="{padding: 0}" class="layout-header-bar">
                    <Icon @click.native="collapsedSider" :class="rotateIcon" :style="{margin: '0 20px'}" type="md-menu" size="24"></Icon>
                    <Dropdown style="margin-right: 50px; float: right" trigger="click">
                        <Avatar style="background-color: #87d068" icon="ios-person" />
                        <DropdownMenu slot="list">
                            <DropdownItem>设置</DropdownItem>
                            <DropdownItem>注销</DropdownItem>
                        </DropdownMenu>
                    </Dropdown>
                </Header>

                <Content :style="{padding: '0 16px 16px'}">
                    <Breadcrumb :style="{margin: '16px 0'}">
                        <BreadcrumbItem>Home</BreadcrumbItem>
                        <BreadcrumbItem>Components</BreadcrumbItem>
                        <BreadcrumbItem>Layout</BreadcrumbItem>
                    </Breadcrumb>
                    <Card shadow="true" dis-hover="true">
                        <keep-alive>
                            <YTable></YTable>
                        </keep-alive>
                    </Card>
                </Content>
            </Layout>
        </Layout>
    </div>
</template>

<script>

    import YTable from './subs/ytable/YTable'

    export default {
        name: "Main",
        data () {
            return {
                isCollapsed: false
            }
        },
        components:{
            YTable
        },
        computed: {
            rotateIcon () {
                return [
                    'menu-icon',
                    this.isCollapsed ? 'rotate-icon' : ''
                ];
            },
            menuitemClasses () {
                return [
                    'menu-item',
                    this.isCollapsed ? 'collapsed-menu' : ''
                ]
            }
        },
        methods: {
            collapsedSider () {
                this.$refs.side1.toggleCollapse();
            }
        }
    }
</script>

<style scoped lang="scss">
    .layout{
        border: 1px solid #d7dde4;
        background: #f5f7f9;
        position: relative;
        border-radius: 4px;
        overflow: hidden;
        height: 100%;
        .ivu-layout{
            height: 100%;
        }
        h1{
            color: white;
            text-align: center;
            padding: 10px 0;
        }
    }
    .layout-header-bar{
        background: #fff;
        box-shadow: 0 1px 1px rgba(0,0,0,.1);
    }
    .menu-item span{
        display: inline-block;
        overflow: hidden;
        width: 69px;
        text-overflow: ellipsis;
        white-space: nowrap;
        vertical-align: bottom;
        transition: width .2s ease .2s;
    }
    .menu-item i{
        transform: translateX(0px);
        transition: font-size .2s ease, transform .2s ease;
        vertical-align: middle;
        font-size: 16px;
    }
    .collapsed-menu span{
        width: 0px;
        transition: width .2s ease;
    }
    .collapsed-menu i{
        transform: translateX(5px);
        transition: font-size .2s ease .2s, transform .2s ease .2s;
        vertical-align: middle;
    }
</style>
