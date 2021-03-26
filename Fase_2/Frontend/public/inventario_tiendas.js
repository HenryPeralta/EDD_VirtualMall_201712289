var elemento = ""
var contador = 0
var depa = ""
var info = ""
var calificacion = 0
var List_productos = [];

fetch('http://localhost:3000/cargaproduct')
.then(response => response.json())
.then(data => {
    //elemento.innerHTML = `<p>${data.Datos[0].Departamentos[0].Tiendas[0].Nombre}</p>`;
    for(var i =0; i < (data.Datos).length; i++){
        var departamentos = data.Datos[i].Departamentos
        //console.log(departamentos)
        for(var j=0; j < departamentos.length; j++){
            tiendas = departamentos[j].Tiendas
            depa = departamentos[j].Nombre
            //console.log(tiendas)
            for(var k=0; k < tiendas.length; k++){
                info = tiendas[k].Nombre
                var image = tiendas[k].Logo
                var descrip = tiendas[k].Descripcion
                var contacto = tiendas[k].Contacto
                calificacion = tiendas[k].Calificacion

                //var calificacion = tiendas[k].Calificacion
                //contador += 1
                //console.log(info)
                elemento += `<div id='${contador += 1}' class='div_tienda'>
                                <center><h1>Tienda</h1><p>Nombre: ${info} <br> Descripcion: ${descrip} <br> Contacto: ${contacto}</p>                               
                                <img src='${image}' class='image'/> <br>  
                                </center>
                                <br>
                                </div>
                                <br><br>
                                <STYLE type="text/css"> 
                                .div_tienda{width: 500px; margin:0px auto; border: #333 2px solid;} 
                                .image{width: 300px; height: 250px;}
                                </STYLE>`
            }
        }
    }
    //console.log(data)
    document.getElementById("tiendas").innerHTML = elemento;
})

function productos(nombre_tienda, nombre_departamento, no_calificacion, nombre_p, codigo_p, descripcion_p, precio_p, cantidad_p, imagen_p){
    this.nombre_tienda = nombre_tienda;
    this.nombre_departamento = nombre_departamento;
    this.no_calificacion = no_calificacion;
    this.nombre_p = nombre_p;
    this.codigo_p = codigo_p;
    this.descripcion_p = descripcion_p;
    this.precio_p = precio_p;
    this.cantidad_p = cantidad_p;
    this.imagen_p = imagen_p;
}

function mandar_producto(){
    var contenido = document.getElementById('agregando_product').value;
    fetch('http://localhost:3000/producto', {
        method: 'POST',
        headers:{
            'Content-Type': 'application/json'
        },
        body:contenido
    })

    fetch('http://localhost:3000/producto')
    .then(response => response.json())
    .then(data =>{
        for(var i =0; i < (data.Productos).length; i++){
            var product_n = data.Productos[i].Nombre
            var product_codigo = data.Productos[i].Codigo
            var product_descripcion = data.Productos[i].Descripcion
            var product_precio = data.Productos[i].Precio
            var product_cantidad = data.Productos[i].Cantidad
            var product_imagen = data.Productos[i].Imagen
            var local = new productos(info, depa, calificacion, product_n, product_codigo, product_descripcion, product_precio, product_cantidad, product_imagen);
            List_productos.push(local);
        }
    })

    fetch('http://localhost:3000/Inventarios2')
    .then(responses => responses.json())
    .then(datax =>{
        for(var a = 0; a < (datax.Invetarios).length; a++){
            var nombre_tienda = datax.Invetarios[a].Tienda
            var departamentox = datax.Invetarios[a].Departamento
            var calificacion = datax.Invetarios[a].Calificacion
            var tiendas = datax.Invetarios[a].Productos
            for(var b =0; b < tiendas.length; b++){
                var product_n = tiendas[b].Nombre
                var product_codigo = tiendas[b].Codigo
                var product_descripcion = tiendas[b].Descripcion
                var product_precio = tiendas[b].Precio
                var product_cantidad = tiendas[b].Cantidad
                var product_imagen = tiendas[b].Imagen
                var locales = new productos(nombre_tienda, departamentox, calificacion, product_n, product_codigo, product_descripcion, product_precio, product_cantidad, product_imagen);
                List_productos.push(locales);
            }
        }
    })
    mandar();
    mostrarlista();
}

function mostrarlista(){
    var lista ='';
    for(var i=0; i < List_productos.length; i++){
        lista += 'nombre tienda: ' + List_productos[i].nombre_tienda +
        ' departamento: ' + List_productos[i].nombre_departamento +
        ' calificacion: ' + List_productos[i].no_calificacion + 
        ' nombre producto: ' + List_productos[i].nombre_p +
        ' codigo: ' + List_productos[i].codigo_p +
        ' descripcion: ' + List_productos[i].descripcion_p +
        ' precio: ' + List_productos[i].precio_p +
        ' cantidad: ' + List_productos[i].cantidad_p + 
        ' imagen: ' + List_productos[i].imagen_p + '\n';
    }
    console.log(lista);
}

function mandar(){
    /*var datos ='';
    for(var i=0; i < List_productos.length; i++){
        datos += "Tienda:" + List_productos[i].nombre_tienda +
        "Departamento:"  + List_productos[i].nombre_departamento +
        "Calificacion:" + List_productos[i].no_calificacion +
        "Productos: [{" + 
        "Nombre:"  + List_productos[i].nombre_p +
        "Codigo:" + List_productos[i].codigo_p +
        "Descripcion:" + List_productos[i].descripcion_p +
        "Precio:" + List_productos[i].precio_p +
        "Cantidad:" + List_productos[i].cantidad_p + 
        "Imagen:" + List_productos[i].imagen_p + 
        "]}";
    }*/
    fetch('http://localhost:3000/Inventarios2', {
        method: 'POST',
        headers:{
            'Content-Type': 'application/json'
        },
        //body: JSON.stringify({"Invetarios":[{"Tienda": List_productos[i].nombre_tienda, "Departamento": List_productos[i].nombre_departamento, "Calificacion": List_productos[i].no_calificacion, "Productos":[{"Nombre": List_productos[i].nombre_p, "Codigo": List_productos[i].codigo_p, "Descripcion": List_productos[i].descripcion_p, "Precio": List_productos[i].precio_p, "Cantidad": List_productos[i].cantidad_p, "Imagen": List_productos[i].imagen_p}]}]})
        body: JSON.stringify({"Inventarios":[{datos}]})
    })
}

/*http.post('http://localhost:3000/Inventarios2',function(req, res){
    for(var i =0; i < (req.body.Productos).length; i++){
        var product_n = data.Productos[i].Nombre
        var product_codigo = data.Productos[i].Codigo
        var product_descripcion = data.Productos[i].Descripcion
        var product_precio = data.Productos[i].Precio
        var product_cantidad = data.Productos[i].Cantidad
        var product_imagen = data.Productos[i].Imagen
        var local = new productos(info, depa, calificacion, product_n, product_codigo, product_descripcion, product_precio, product_cantidad, product_imagen);
        List_productos.push(local);
    }
    res.json( {...req.body, Tienda: analizadorLexico.lista_de_tokens} );
});*/