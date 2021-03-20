<template>
    <el-col :xl="6" :lg="6" :md="8" :sm="12" :xs="24" class="card-gutter">
        <a class="card-url" :href="cardData.url">
            <el-card class="card-cursor card-border">
                <div slot="header" class="clearfix card-header font-color-navyblue">
                    <span>{{cardData.name}}</span>
                </div>
                <div class="card-desc font-color-black">{{cardData.desc}}</div>
            </el-card>
        </a>
    </el-col>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
    name: 'pageCard',
    props: ['data'],
    data() {
        return {
            cardData: this.data
        };
    }
});
</script>

<style lang="scss">
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

.clearfix:before,
.clearfix:after {
    display: table;
    content: '';
}
.clearfix:after {
    clear: both;
}

.card-header {
    font-size: 20px;
    font-weight: bold;
}

.card-desc {
    font-size: 18px;
}
</style>
