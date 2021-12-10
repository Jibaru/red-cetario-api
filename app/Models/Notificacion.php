<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Notificacion extends Model
{
    use HasFactory;

    protected $table = 'comentarios';

    protected $fillable =
        [
            'id',
            'titulo',
            'descripcion',
            'fecha_envio',
            'fecha_visto',
            'id_cliente',
            'created_at',
            'updated_at'
        ];
}
