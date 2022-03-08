## Отображение графов

# Содержание
- [Содержание](#содержание)
  - [Описание](#описание)
  - [Использование](#использование)
  - [Новые задачи](#задачи)
  
## Описание
Простая программа для рисования графов на основе определенных входных данных таких как:
- матрица инцидентности

для рисования графов на языке Go используется библиотека:
- go-graphviz

## Использование
Для того, чтобы создать определенный граф на основе матрицы инцидентности необходимо вводить
матрицу таким образом:

Пример №1:
-1  1
 1 -1

Пример №2:
 1 -1
-1  0
 0  0 

значения вводятся через пробел. В конце строки необходимо внести enter.

Значения, которые принимаются : 
-1 - это начало ребра. 
 1 - это конец ребра.
 0 - это без ребер.

После того, как ввели матрицу нажать в кнопку граф для формирования и отображения графа. 
