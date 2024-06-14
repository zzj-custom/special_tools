<script setup lang="ts">
import { computed, ref } from 'vue';
import type { NeList } from '@/types/netease/index';
const search = ref('');
const props = defineProps({
    title: {
        type: String,
        default: '网易云转换文件列表'
    },
    list: {
        type: Array as () => NeList[],
        default: []
    }
});

const duration = computed(() => {
    {
        return (seconds: number) => {
            const hours = Math.floor(seconds / 3600);
            const minutes = Math.floor((seconds % 3600) / 60);
            const secs = seconds % 60;

            const formattedHours = hours.toString().padStart(2, '0');
            const formattedMinutes = minutes.toString().padStart(2, '0');
            const formattedSecs = secs.toString().padStart(2, '0');

            if (hours > 0) {
                return `${formattedHours}:${formattedMinutes}:${formattedSecs}`;
            }
            return `${formattedMinutes}:${formattedSecs}`;
        };
    }
});
</script>

<script lang="ts">
export default {
    name: 'NeListComponent'
};
</script>

<template>
    <v-card hover>
        <v-card-title class="d-flex align-center pe-2">
            <v-icon color="primary" icon="mdi-video"></v-icon> &nbsp; {{ props.title }}

            <v-spacer></v-spacer>

            <v-text-field
                v-model="search"
                density="compact"
                label="Search"
                prepend-inner-icon="mdi-magnify"
                variant="solo-filled"
                flat
                hide-details
                single-line
            ></v-text-field>
        </v-card-title>

        <v-divider></v-divider>

        <v-data-table v-model:search="search" :items="props.list">
            <template v-slot:item.albumPic="{ item }">
                <v-card class="my-2" elevation="2" rounded>
                    <v-img :src="item.albumPic" height="64" cover></v-img>
                </v-card>
            </template>

            <template v-slot:item.duration="{ item }">
                <p>{{ duration(item.duration) }}</p>
            </template>

            <template v-slot:item.format="{ item }">
                <div class="text-end">
                    <v-chip
                        :color="item.format ? 'green' : 'red'"
                        :text="item.format ? 'In stock' : 'Out of stock'"
                        class="text-uppercase"
                        size="small"
                        label
                    ></v-chip>
                </div>
            </template>
        </v-data-table>
    </v-card>
</template>
