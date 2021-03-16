fetch('http://localhost:3000/Tiendas')
.then(response => response.json())
.then(data => {
    var elemento = ""
    //elemento.innerHTML = `<p>${data.Datos[0].Departamentos[0].Tiendas[0].Nombre}</p>`;
    for(var i =0; i < (data.Datos).length; i++){
        var departamentos = data.Datos[i].Departamentos
        //console.log(departamentos)
        for(var j=0; j < departamentos.length; j++){
            var tiendas = departamentos[j].Tiendas
            //console.log(tiendas)
            for(var k=0; k < tiendas.length; k++){
                var info = tiendas[k].Nombre
                var image = tiendas[k].Logo
                //console.log(info)
                elemento += `<p>${info}</p> <br> <img src='${image}' height = '${300}'/> <br>`
            }
        }
    }
    //console.log(data)
    document.getElementById("tienda").innerHTML = elemento;
})
