<template>
    <el-container class="navpage">
        <el-header height="120px">
            <div class="navpage-header">
                <div class="navpage-logo">
                    <img src="../assets/logo.png" class="navpage-logo-img" alt="WindSpirit IT Doc Lib" />
                </div>
                <h1 class="navpage-title font-color-white">WindSpirit IT Doc Lib</h1>
            </div>
        </el-header>
        <el-main style="padding-top: 0">
            <div class="navpage-desc font-color-white">
                <p style="font-size: 60px;font-weight: bold;">
                    Document<br />
                    Everything!
                </p>
                <p style="font-size: 25px;font-weight: lighter;margin-block-start: 0">
                    Used for storage study notes<br />
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
                        <div v-for="item in pages" :key="item.id">
                            <page-card :cardData="item"></page-card>
                        </div>
                    </div>
                </el-row>
            </div>
        </el-main>
    </el-container>
</template>

<script lang="ts">
import Vue from 'vue';
import Api from '@/utils/api';
import pageCard from '@/components/pageCard.vue';

const api = new Api();

export default Vue.extend({
    name: 'NavPage',
    components: {
        pageCard
    },
    data() {
        return {
            cardInfoGroupByCat: {}
        };
    },
    mounted() {
        let _this = this;
        _this.getPagesInfo();
    },
    methods: {
        getPagesInfo() {
            let _this = this;
            api.pages()
                .then((res) => {
                    let cardInfo = {};
                    let pagesInfo = res.data.pages;
                    Object.keys(pagesInfo).forEach((key) => {
                        let cat = pagesInfo[key].cat;
                        if (!Object.prototype.hasOwnProperty.call(cardInfo, cat)) {
                            cardInfo[cat] = [];
                        }
                        cardInfo[cat].push(pagesInfo[key]);
                    });
                    _this.cardInfoGroupByCat = cardInfo;
                    _this.$forceUpdate();
                })
                .catch((e) => {
                    console.log(e);
                });
        }
    }
});
</script>

<style lang="scss">
.navpage {
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    background-image: url('../assets/bg.png');
    background-position: 50% 50%;
    background-size: cover;
    background-repeat: no-repeat;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    min-height: 900px;
    padding: 0 calc((100% - 1200px) / 2);
}

.font-color-white {
    color: #ffffff;
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
</style>
