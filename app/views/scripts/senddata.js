// textareaのデータをhttpで送る

// inputのデータを取得
// 取得したデータをサーバーへリクエスト
function createNewLine() {

      let url = "https://weather.tsukumijima.net/api/forecast";
      fetch(url + "?city=410020").then(function(response) {
            console.log(response);
            if (response.ok) {
                  response.json().then(function(json) {
                        console.log(json.publicTime);
                        outputId.value = json.publicTime;
                  })
                  .catch(e => {
                        console.log(e.message);
                  });
            } else {
                  outputId.value = "fatal";
            }
      });
}