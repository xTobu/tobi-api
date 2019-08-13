# tobi-api
## Develop
### Golang live reloading
```bash
realize start --run
```

## 資料夾結構
- ### lib
- ### utils
- ### helpers

## 群益
``` python
# 5-6 SKSTOCK ( 國內報價商品物件 )
struct SKSTOCK
{
    SHORT sStockidx; # 商品自定索引代號
    SHORT sDecimal; # 小數位數
    SHORT sTypeNo;  
    BSTR bstrMarketNo; # 市埸代碼
    BSTR bstrStockNo; # 商品代碼
    BSTR bstrStockName; # 商品名稱
 
    LONG nHigh; #  最高價
    LONG nOpen; #  開盤價
    LONG nLow; #  最低價
    LONG nClose; #  成交價
 
    LONG nTickQty; #  單量
 
    LONG nRef; #  昨收、參考價
 
    LONG nBid; #  買價
    LONG nBc;  #  買量
    LONG nAsk; #  賣價
    LONG nAc;  #  賣量
 
    LONG nTBc; #  買盤量
    LONG nTAc; #  賣盤量
 
    LONG nFutureOI; #  期貨未平倉 OI
 
    LONG nTQty; #  總量
    LONG nYQty; #  昨量
 
    LONG nUp; #  漲停價
    LONG nDown; #  跌停價
    LONG nSimulate; # 揭示 0:一般 1:試算
};
```