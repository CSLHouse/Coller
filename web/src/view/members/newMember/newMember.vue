<template>
    <el-menu class="el-menu-demo" mode="horizontal">
      <el-menu-item index="1">新会员办理</el-menu-item>
    </el-menu>
    <div class="h-6" />
    <el-form
      ref="ruleFormRef"
      :model="memberForm"
      status-icon
      label-width="120px"
      class="demo-ruleForm"
      style="max-width: 460px"
    >
      <el-form-item label="会员卡号" prop="cardId">
        <el-input v-model.number="memberForm.cardId" type="number" autocomplete="off" />
      </el-form-item>
      <el-form-item label="联系电话" prop="telephone">
        <el-input v-model.number="memberForm.telephone" type="tel" autocomplete="off" />
      </el-form-item>
      <el-form-item label="姓名" prop="name">
        <el-input v-model="memberForm.memberName" type="text" autocomplete="off" />
      </el-form-item>
      <el-form-item label="购买会员卡" prop="comboId">
        <el-select v-model="comboOption" class="m-2" placeholder="请选择会员卡" size="large">
            <el-option
            v-for="item in comboOptions"
            :key="item.key"
            :label="item.value"
            :value="item"
            />
        </el-select>
      </el-form-item>
      <el-form-item label="赠送" prop="gift">
        <el-col :span="8">
          <el-input v-model.number="memberForm.remainTimes" type="number" autocomplete="off" />
        </el-col>
        <a style="color: red; margin-left: 10px;">整数，填写赠送的天数/次数/金额</a>
      </el-form-item>
      <el-form-item label="实收金额" prop="collection">
        <el-col :span="8">
          <el-input v-model.number="memberForm.collection" type="number" autocomplete="off" />
        </el-col>
        <a style="color: red; margin-left: 10px;">整数，必须与选择的Vip卡价格一致</a>
      </el-form-item>
      <el-form-item label="开始日期" prop="date">
        <div class="demo-date-picker">
          <el-date-picker
            v-model="memberForm.startDate"
            type="date"
            format="YYYY/MM/DD"
            value-format="YYYY-MM-DD"
          >
          </el-date-picker>
        </div>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm(ruleFormRef)"
          >确认</el-button
        >
        <el-button @click="resetForm(ruleFormRef)">清空</el-button>
      </el-form-item>
    </el-form>
  </template>
  
  <script lang="ts" setup>
  import { createVIPMember } from '@/api/member'
  import { reactive, ref, onBeforeMount, watch } from 'vue'
  import { FormInstance, FormRules, ElMessage } from 'element-plus'
  import { Plus } from '@element-plus/icons-vue'
  import { comboStore } from '@/pinia/modules/combo'
  
  const ruleFormRef = ref<FormInstance>()
  
  const validateCardId = (rule: any, value: any, callback: any) => {
    if (value === '') {
      callback(new Error('Please input the id'))
    } else {
      if (memberForm.cardId !== '') {
        if (!ruleFormRef.value) return
        ruleFormRef.value.validateField('cardId', () => null)
      }
      callback()
    }
  }
  
  const validateTelephone = (rule: any, value: any, callback: any) => {
    if (value === '') {
      callback(new Error('Please input the phone'))
    } else {
      if (memberForm.telephone !== '') {
        if (!ruleFormRef.value) return
        ruleFormRef.value.validateField('telephone', () => null)
      }
      callback()
    }
  }

  const validateAmount = (rule: any, value: any, callback: any) => {
    if (value === '') {
      callback(new Error('Please input the amount'))
    } else {
      if (memberForm.collection <= 0) {
        if (!ruleFormRef.value) return
        ruleFormRef.value.validateField('collection', () => null)
      }
      callback()
    }
  }

  // 查询会员卡类型
  const vipComboStore = comboStore()
  let comboList = []
  interface comboItem {
    key: number,
    value: string,
    price: number
  }
  const comboOption = ref<comboItem>()
  const comboOptions = ref<comboItem[]>([])
 
  const getComboData = async() => {
    comboList = await vipComboStore.comboList
    if (comboList['length'] < 1) {
      await vipComboStore.GetAllVIPCombos()
      comboList = vipComboStore.comboList
    }
    comboOptions.value = comboList.map((item) => {
      return {key: item.comboId, value: item.comboName, price: item.comboPrice}
    })
  }
  getComboData()
  
  watch(() => comboOption.value, () => {
    memberForm.comboId = comboOption.value.key
    memberForm.comboName = comboOption.value.value
    console.log("--------watch------memberForm:", memberForm)
  })

  const memberForm = reactive({
    cardId: null,
    telephone: null,
    memberName: '',
    comboId: 0,
    comboName: '',
    remainTimes: null,
    startDate: '',
    collection: null,
  })

  // const rules = reactive<FormRules<typeof memberForm>>({
  //   cardId: [{ validator: validateCardId, trigger: 'blur' }],
  //   telephone: [{ validator: validateTelephone, trigger: 'blur' }],
  //   collection: [{ validator: validateAmount, trigger: 'blur' }],
  // })

  function nowDate() {
    let time = new Date()
    let y = time.getFullYear()
    let m: any = time.getMonth() + 1
    let d: any = time.getDate()
    if (m < 10) m = '0' + m
    if (d < 10) d = '0' + d
    const date = `${y}-${m}-${d}`
		return date
	};

  onBeforeMount(() => {
    memberForm.startDate = nowDate()
  })

  const submitForm = async(formEl: FormInstance | undefined) => {
    if (!formEl) return false
    if (memberForm.cardId < 1 || memberForm.telephone == '' 
    || memberForm.comboId < 1 || memberForm.collection != comboOption.value.price ) {
      ElMessage({
        type: 'error',
        message: '填写信息有误'
      })
      return
    }
    let res
    formEl.validate(async(valid) => {
      if (valid) {
        res = await createVIPMember(memberForm)
        
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '添加成功'
          })
          formEl.resetFields()
          memberForm.cardId = null
          memberForm.telephone = null
          memberForm.memberName = ''
          memberForm.comboId = null
          memberForm.comboName = ''
          memberForm.remainTimes = null
          memberForm.collection = null
          comboOption.value = {key: 0, value: "", price: 0}
        }
      } else {
        return false
      }
      
    })
    
  }
  
  const resetForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    comboOption.value = {key: 0, value: "", price: 0}
    formEl.resetFields()
    memberForm.cardId = null
    memberForm.telephone = null
    memberForm.memberName = ''
    memberForm.comboId = null
    memberForm.comboName = ''
    memberForm.remainTimes = null
    memberForm.collection = null
  }
  
  </script>
