<template>
  <div class="main">
    <div class="card-body border-bottom py-2">
      <div class="d-flex align-items-center">
        <div class="ms-auto">
          <div class="input-icon">
            <span class="input-icon-addon">
              <IconSearch class="icon" />
            </span>
            <input
              type="text"
              v-model="searchQuery"
              class="form-control form-control-rounded"
              placeholder="Busca..."
            />
          </div>
        </div>
      </div>
    </div>
    <div class="card">
      <div class="table-responsive">
        <table class="table table-sm table-bordered table-vcenter table-hover card-table">
          <thead>
            <tr>
              <th
                v-for="column in columns"
                :key="column.field"
                @click="sortBy(column.field)"
                class="text-nowrap cursor-pointer text-start"
                :class="{ 'text-primary': sortColumn === column.field }"
              >
                {{ column.title }}
                <span v-if="sortColumn === column.field" class="ms-1">
                  {{ sortOrder === 'asc' ? '↑' : '↓' }}
                </span>
              </th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="row in paginatedData"
              :key="row.id"
              class="clickable-row"
              @click="ruta(row.Nombre)"
            >
              <td v-for="column in columns" :key="column.field" class="text-start fw-medium">
                {{ row[column.field] }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { router } from '@router/router'
import { IconSearch } from '@tabler/icons-vue'
import { computed, ref } from 'vue'

interface Column {
  field: string
  title: string
}
interface Row {
  [key: string]: any
}

const props = defineProps<{
  rows: Row[]
  columns: Column[]
  title?: string
  initialPageSize?: number
}>()

const currentPage = ref(1)
const sortColumn = ref('')
const sortOrder = ref<'asc' | 'desc'>('asc')
const searchQuery = ref('')
const pageSize = ref(props.initialPageSize || 10)

const filteredRows = computed(() => {
  return props.rows.filter((row) =>
    Object.values(row).some((value) =>
      String(value).toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  )
})

const paginatedData = computed(() => {
  let sortedData = [...filteredRows.value]
  if (sortColumn.value) {
    sortedData.sort((a, b) => {
      if (a[sortColumn.value] < b[sortColumn.value]) return sortOrder.value === 'asc' ? -1 : 1
      if (a[sortColumn.value] > b[sortColumn.value]) return sortOrder.value === 'asc' ? 1 : -1
      return 0
    })
  }
  const start = (currentPage.value - 1) * pageSize.value
  const end = start + pageSize.value
  return sortedData.slice(start, end)
})

const sortBy = (column: string) => {
  if (sortColumn.value === column) {
    sortOrder.value = sortOrder.value === 'asc' ? 'desc' : 'asc'
  } else {
    sortColumn.value = column
    sortOrder.value = 'asc'
  }
}

const ruta = async (area: string) => {
  await router.push({
    name: 'inofarea',
    params: {
      area: area
    }
  })
}
</script>

<style lang="scss" scoped>
.main {
  width: min-content;
  .card {
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    border-radius: 0.5rem;
    width: max-content;
  }

  .card-header {
    background-color: #f8fafc;
    border-bottom: 1px solid #e2e8f0;
  }

  .input-icon {
    position: relative;

    .input-icon-addon {
      position: absolute;
      left: 0.75rem;
      top: 50%;
      transform: translateY(-50%);
      z-index: 2;
      color: #888;
    }

    input {
      padding-left: 2.5rem;
      border-radius: 2rem;
      border: 1px solid #e2e8f0;
      transition: border-color 0.3s ease;

      &:focus {
        border-color: var(--tblr-primary);
        box-shadow: 0 0 0 0.2rem rgba(var(--tblr-primary-rgb), 0.25);
      }
    }
  }

  .table {
    width: min-content;
    thead {
      th {
        font-size: 0.76rem;
        font-weight: 500;
        color: #6c757d;
        padding: 0.5rem 0.75rem;
        text-transform: none;
        border-bottom: 1px solid #e2e8f0;
      }
      :first-child {
        min-width: 280px !important;
      }
      :last-child {
        width: min-content;
      }
    }

    tbody {
      tr.clickable-row {
        transition: background-color 0.3s ease;

        &:hover {
          background-color: rgba(var(--tblr-primary-rgb), 0.1);
          cursor: pointer;
        }

        td {
          font-size: 0.7rem;
          padding: 0.5rem 0.75rem;
          color: #343a40;
          border-bottom: 1px solid #e2e8f0;
        }
      }
    }
  }
}
</style>
