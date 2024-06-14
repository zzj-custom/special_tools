<template>
    <NeUpload @list="handleList"></NeUpload>
    <NeListComponent :list="list"></NeListComponent>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import NeListComponent from '@/components/netease/list.vue';
import NeUpload from '@/components/netease/upload.vue';
import { NeList } from '@/types/netease/index';

const list = ref<NeList[]>([]);

const handleList = (l: NeList[]) => {
    console.log('当前接收到的子组件数据', l);
    list.value = l;
};
</script>
