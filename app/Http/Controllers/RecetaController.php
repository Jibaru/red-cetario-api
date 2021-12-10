<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Models\Receta;

class RecetaController extends Controller
{
    public function index() 
    {
        $recetas = Receta::all();

        return array(
            "ok" => count($recetas) > 0,
            "recetas" => $recetas
        );
    }

    public function show($id)
    {
        $receta = Receta::with(['comentarios', 'ingredientes', 'materiales', 'pasos'])
            ->get()
            ->find($id);
        return $receta;
    }
}
