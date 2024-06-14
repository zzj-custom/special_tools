<script setup lang="ts">
import { ref } from 'vue';
import { netcase, response } from '../../../wailsjs/go/models';
import { ParseInfo } from '../../../wailsjs/go/App/Netcase';
import { NeList } from '@/types/netease';
import { emit } from 'process';

const input = ref();

const fileList = ref<NeList[]>([]);

const emit = defineEmits(['list']);

const triggerInput = () => {
    if (!input.value) {
        console.log('input属性不存在');
    }
    input.value?.click();
};

const change = (e: any) => {
    const file = e.target.files;
    if (file.length === 0) {
        fileList.value = [];
        return;
    }

    fileList.value = [];
    const filesData: netcase.ParseInfoDTO[] = [];
    // 遍历文件列表，并将文件对象添加到数组中
    for (const f of file) {
        console.log('处理文件：', f, '文件名：', f.name);
        const reader = new FileReader();
        reader.onload = (e) => {
            console.log('开始处理数据');
            const fileContent = e.target?.result;
            if (typeof fileContent !== 'string') {
                console.error('File content is not a string:');
            } else {
                filesData.push({
                    name: f.name,
                    content: fileContent
                });
            }
            // 检查是否已读取所有文件
            processFiles(filesData);
        };
        reader.readAsDataURL(f);
    }
};

const processFiles = async (p: netcase.ParseInfoDTO[]) => {
    try {
        const response: response.Reply = await ParseInfo(p);
        console.log('response:', response);
        response.result.forEach((item: any) => {
            fileList.value.push({
                name: item.name,
                musicID: item.musicID,
                musicName: item.musicName,
                albumPic: item.albumPic,
                duration: item.duration,
                format: item.format
            });
            emit('list', fileList.value);
        });
    } catch (error) {
        console.error('Error uploading files:', error);
    }
};
</script>

<script lang="ts">
export default {
    name: 'NeUpload'
};
</script>

<template>
    <v-card class="mx-auto" hover>
        <v-card-item>
            <v-card-title> 文件上传 </v-card-title>

            <v-card-subtitle> 选择你需要转换的网易云文件 </v-card-subtitle>
        </v-card-item>

        <v-card-text class="text-center">
            <div>
                <v-row justify="center">
                    <v-col cols="auto">
                        <v-img
                            src="https://cdn.vuetifyjs.com/images/parallax/material.jpg"
                            alt="Profile Image"
                            width="120"
                            height="120"
                            cover
                            class="rounded-circle"
                        ></v-img>
                    </v-col>
                </v-row>

                <div class="d-flex justify-center my-4 gap-3">
                    <v-btn class="bg-primary" @click="triggerInput">Upload </v-btn>
                    <v-file-input multiple ref="input" style="display: none" @change="change"></v-file-input>
                    <v-hover v-slot="{ isHovering, props }"
                        ><v-btn
                            v-bind="props"
                            :class="{ 'on-hover': isHovering, 'bg-error': isHovering }"
                            class="text-none"
                            :color="isHovering ? 'undeinfed' : 'error'"
                            variant="outlined"
                            >Reset</v-btn
                        >
                    </v-hover>
                </div>
                <p class="mb-0">Allowed JPG, GIF or PNG. Max size of 800K</p>
            </div>
        </v-card-text>
    </v-card>
</template>

<style scoped lang="scss"></style>
