<template>
  <el-container>
      <el-header>
          Header
      </el-header>
      <el-main>
          <el-row>
              <el-col :xl="4" :lg="6" :md="8" :sm="12" :xs="24" class="card-gutter"
                      v-for="item in test" :key="item">
                  <el-card shadow="hover">
                      <div slot="header" class="clearfix">
                          <span>Title {{item}}</span>
                      </div>
                      <div>Descriptions {{item}}</div>
                  </el-card>
              </el-col>
          </el-row>
      </el-main>
  </el-container>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
    name: 'NavPage',
    data () {
        return {
            test: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
        }
    }
})
</script>

<style>
    .card-gutter {
        padding: 0 10px;
        margin-bottom: 10px;
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
