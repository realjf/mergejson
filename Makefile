# ##############################################################################
# # File: Makefile                                                             #
# # Project: mergejson                                                         #
# # Created Date: 2023/11/07 18:14:44                                          #
# # Author: realjf                                                             #
# # -----                                                                      #
# # Last Modified: 2023/11/07 18:15:32                                         #
# # Modified By: realjf                                                        #
# # -----                                                                      #
# #                                                                            #
# ##############################################################################


B ?= master
M ?= "update"

.PHONY: push
push:
	@git add -A && git commit -m ${M} && git push origin ${B}
