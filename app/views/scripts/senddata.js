// textareaのデータをhttpで送る

// inputのデータを取得
// 取得したデータをサーバーへリクエスト
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
      const url = '/newline/';
      let data = {text: inputId.value};
      getNewLine(url, data).then(data => {
            console.log(data);
            outputId.value = data.text;
      })
}