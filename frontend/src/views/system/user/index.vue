<script setup lang="ts">
defineProps<{
  embedded?: boolean
}>()
</script>

<template>
  <n-card title="用户管理">
    <n-space vertical size="large">
      <n-alert type="info" title="工程骨架">
        多用户架构入口，后续支持用户新增、重置密码、分配角色、绑定租户。
      </n-alert>
      <n-data-table :columns="[]" :data="[]" />
    </n-space>
  </n-card>
</template>
