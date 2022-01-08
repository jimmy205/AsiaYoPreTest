### 匯率轉換API

#### 格式
```
Method: POST
Path  : api/currency/transfer

request {
    from   string
    to     string
    amount float64
}

response {
    amount string
}
```

#### 服務啟動方式
```
go run .

可帶入port號參數: PORT (預設8000) ex: PORT=8000 go run .
```

#### 資料夾結構

##### router
路由層
```
1. 放置跟http服務有關的資料夾(目前採用gin)
2. 設定API路徑，API路徑依功能依據分開 (currency ...)
```

##### api
入口層
```
1. API呼叫入口
2. 資料夾依功能分(currencyApi...)
3. 驗證方式採用gin validator
4. 錯誤於正式專案需另外處理 ex. 參數未填,錯誤代碼,錯誤訊息...等。目前先將錯誤一律傳出。
5. 可把此層當成controller看待，將會在此層呼叫所需的BIN後將資料組合傳出
```

##### business
商業邏輯層
```
1. 撰寫商業邏輯的地方
2. 通常會於此層呼叫DB、外接API拿取所需資料並組裝回API層
```

##### repository
資料庫、第三方層
```
1. DB實作
2. 第三方實作
```

##### 其它
```
1. 因測驗所以沒有分那麼多層，目前先把對照表寫死在 business.currency 中
2. 對照表應是從DB或是第三方取得，可在對照表取得地方加上快取
3. 因測試性質關係，寫的較為簡陋。如有問題歡迎提問
```

