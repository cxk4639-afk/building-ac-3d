<script setup lang="ts">
defineProps<{
  embedded?: boolean
}>()
</script>

<template>
  <n-card title="工作台">
    <n-space vertical size="large">
      <n-alert type="info" title="工程骨架">
        后续展示平台运行概览、租户数量、设备数量、配置数量、最近操作。
      </n-alert>
      <n-data-table :columns="[]" :data="[]" />
    </n-space>
  </n-card>
</template>
