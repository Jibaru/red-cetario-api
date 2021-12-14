<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Models\Comentario;
use App\Models\Receta;

class RecetaController extends Controller
{
    public function index() 
    {
        $recetas = Receta::with(['cliente'])->get();

        return array(
            "ok" => count($recetas) > 0,
            "recetas" => $recetas
        );
    }

    public function show($id)
    {
        $receta = Receta::with([
            'comentarios', 
            'ingredientes', 
            'materiales', 
            'pasos', 
            'cliente',
            'clientes_favoritos'
            ])
            ->get()
            ->find($id);
        return $receta;
    }

    public function comentar(Request $request, $id)
    {
        
        $comentario = new Comentario();

        $comentario->descripcion = $request->descripcion;
        $comentario->id_receta = $id;
        $comentario->id_cliente = $request->id_cliente;

        return array(
            "ok" => $comentario->save(),
            "comentario" => $comentario
        );
    }
}
