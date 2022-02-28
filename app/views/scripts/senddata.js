// textareaのデータをhttpで送る

// inputのデータを取得
// 取得したデータをサーバーへリクエスト
// function test() {

//       let url = "http://localhost/newline";
//       fetch(url).then(function(response) {
//             console.log(response);
//             if (response.ok) {
//                   await response.json().then(function(json) {
//                         console.log(json.publicTime);
//                         outputId.value = json.publicTime;
//                   })
//                   .catch(e => {
//                         console.log(e.message);
//                   });
//             } else {
//                   outputId.value = "fatal";
//             }
//       });
// }

async function getNewLine(url = '', data = {}) {
      const response = await fetch(url, {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            redirect: 'error',
            referrerPolicy: 'no-referrer',
            body: JSON.stringify(data)
      })
      return response.json();
}

function changeLine() {
      // console.log(typeof(inputId.value))
      const url = '/newline/';
      console.log(url);
      let data = {text: inputId.value};
      getNewLine(url, data).then(data => {
            console.log(data);
            outputId.value = data.text;
      })
}

// function createNewLine() {

//       let url = "https://weather.tsukumijima.net/api/forecast";
//       fetch(url + "?city=410020").then(function(response) {
//             console.log(response);
//             if (response.ok) {
//                   response.text().then(function(text) {
//                         console.log(text);
//                         outputId.value = text;
//                   })
//                   .catch(e => {
//                         console.log(e.message);
//                   });
//             } else {
//                   outputId.value = "fatal";
//             }
//       });
// }