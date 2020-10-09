<template>
  <el-container>
      <el-header height="120px">
          <div class="navpage-header">
              <div class="navpage-logo">
                  <img src="../assets/logo.png" class="navpage-logo-img" alt="WindSpirit IT Doc Lib">
              </div>
              <h1 class="navpage-title font-color-white">WindSpirit IT Doc Lib</h1>
          </div>
      </el-header>
      <el-main style="padding-top: 0">
          <div class="navpage-desc font-color-white">
              <p style="font-size: 60px;font-weight: bold;">
                  Document<br>
                  Everything!
              </p>
              <p style="font-size: 25px;font-weight: lighter;margin-block-start: 0">
                  Used for storage study notes<br>
                  and development documents.
              </p>
          </div>
          <div class="book-cards">
              <el-row>
                  <div v-for="(pages, cat) in cardInfoGroupByCat" :key="cat">
                      <el-col>
                          <el-card class="card-cat" :body-style="{padding: '10px 15px'}">
                              <div>{{cat}}</div>
                          </el-card>
                      </el-col>
                      <el-col :xl="6" :lg="6" :md="8" :sm="12" :xs="24" class="card-gutter"
                              v-for="item in pages" :key="item.id">
                          <a class="card-url" :href="item.url">
                              <el-card class="card-cursor card-border">
                                  <div slot="header" class="clearfix card-header font-color-navyblue">
                                      <span>{{item.name}}</span>
                                  </div>
                                  <div class="card-desc font-color-black">{{item.desc}}</div>
                              </el-card>
                          </a>
                      </el-col>
                  </div>
              </el-row>
          </div>
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
            cardInfoGroupByCat: {}
        }
    },
    mounted () {
        this.getPagesInfo()
    },
    methods: {
        getPagesInfo () {
            axios.post('/api/pages').then(response => {
                const cardInfo = {}
                const pagesInfo = response.data.pages
                Object.keys(pagesInfo).forEach(key => {
                    const cat = pagesInfo[key].cat
                    if (!Object.prototype.hasOwnProperty.call(cardInfo, cat)) {
                        cardInfo[cat] = []
                    }
                    cardInfo[cat].push(pagesInfo[key])
                })
                this.cardInfoGroupByCat = cardInfo
                this.$forceUpdate()
            }).catch(e => {
                console.log(e)
            })
        }
    }
})
</script>

<style>
.font-color-white {
    color: #FFFFFF;
}

.font-color-navyblue {
    color: #183055;
}

.font-color-black {
    color: #000000;
}

.navpage-header {
    width: 100%;
}

.navpage-logo {
    float: left;
    height: 120px;
    margin-left: 10px;
    display: flex;
    align-items: center;
}

.navpage-logo-img {
    width: 60px;
    height: 60px;
}

.navpage-title {
    float: left;
    height: 120px;
    line-height: 120px;
    margin: 0 20px;
}

.navpage-desc {
    text-align: left;
    margin-left: 10px;
}

.navpage-desc p {
    margin-block-start: 0.5em;
    margin-block-end: 0.5em;
}

.book-cards {
    margin-top: 6em;
}

.card-cat {
    width: max-content;
    font-size: 18px;
    color: #183055;
    font-weight: bold;
    margin: 20px 10px;
}

.card-gutter {
    padding: 0 10px;
    margin-bottom: 10px;
    transform: perspective(100px) scale(1);
    transition: 0.2s;
}

.card-gutter:hover {
    -webkit-transform: perspective(100px) scale(1.04);
    -moz-transform: perspective(100px) scale(1.04);
    -ms-transform: perspective(100px) scale(1.04);
    -o-transform: perspective(100px) scale(1.04);
    transform: perspective(100px) scale(1.04);
    transition: 0.2s;
}

.card-url {
    text-decoration: unset;
}

.card-header {
    font-size: 20px;
    font-weight: bold;
}

.card-desc {
    font-size: 18px;
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
