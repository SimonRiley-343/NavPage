<template>
  <el-container>
      <el-header>
          <h1>Document everything</h1>
      </el-header>
      <el-main>
          <el-row>
              <el-col :xl="4" :lg="6" :md="8" :sm="12" :xs="24" class="card-gutter"
                      v-for="item in cardInfo" :key="item.id">
                  <el-card shadow="hover" @click.native="openUrl({url : item.url})" class="card-cursor card-border">
                      <div slot="header" class="clearfix card-header">
                          <span>{{item.name}}</span>
                      </div>
                      <div>{{item.desc}}</div>
                  </el-card>
              </el-col>
          </el-row>
      </el-main>
  </el-container>
</template>

<script lang="ts">
import Vue from 'vue'
import axios from 'axios'

export default Vue.extend({
    name: 'NavPage',
    data () {
        return {
            cardInfo: []
        }
    },
    mounted () {
        this.getPagesInfo()
    },
    methods: {
        openUrl ({ url }: { url: string }) {
            window.location.replace(url)
        },
        getPagesInfo () {
            axios.post('/api/pages').then(response => {
                this.cardInfo = response.data.pages
            }).catch(e => {
                console.log(e)
            })
        }
    }
})
</script>

<style>
    .card-gutter {
        padding: 0 10px;
        margin-bottom: 10px;
    }

    .card-cursor {
        cursor: pointer;
    }

    .card-border {
        border-width: 2px !important;
        border-color: #DDD !important;
    }

    .clearfix:before,
    .clearfix:after {
        display: table;
        content: "";
    }
    .clearfix:after {
        clear: both
    }
</style>
