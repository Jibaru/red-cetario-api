<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateRecetasTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('recetas', function (Blueprint $table) {
            $table->id();
            $table->string('titulo');
            $table->string('descripcion');
            $table->integer('tiempo_prep');
            $table->integer('tiempo_coccion');
            $table->string('url_imagen');
            $table->string('tips')->nullable();
            $table->string('calorias');
            $table->string('dificultad');
            $table->string('cocina');
            $table->unsignedBigInteger('id_cliente')->nullable();
            $table->foreign('id_cliente')
                ->references('id')
                ->on('clientes')
                ->onDelete('set null');
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
        Schema::dropIfExists('recetas');
    }
}
