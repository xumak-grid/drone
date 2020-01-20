FROM drone/drone:0.8-alpine
ADD ./bin/exposer /bin/exposer
ADD ./init.sh /bin/init.sh
ENTRYPOINT ["/bin/init.sh"]
CMD ["/bin/init.sh"]
