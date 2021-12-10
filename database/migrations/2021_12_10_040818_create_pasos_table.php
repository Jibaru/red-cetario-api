<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreatePasosTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('pasos', function (Blueprint $table) {
            $table->id();
            $table->integer('numero_orden');
            $table->string('contenido');
            $table->unsignedBigInteger('id_receta');
            $table->foreign('id_receta')
                ->references('id')
                ->on('recetas')
                ->onDelete('cascade');
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('pasos');
    }
}
