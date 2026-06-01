<script setup lang="ts">
defineProps<{
  embedded?: boolean
}>()
</script>

<template>
  <n-card title="租户管理">
    <n-space vertical size="large">
      <n-alert type="info" title="工程骨架">
        多租户架构入口，后续支持租户新增、停用、套餐、数据隔离配置。
      </n-alert>
      <n-data-table :columns="[]" :data="[]" />
    </n-space>
  </n-card>
</template>
